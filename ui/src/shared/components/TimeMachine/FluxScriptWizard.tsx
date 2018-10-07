// Libraries
import React, {PureComponent} from 'react'

// Components
import {
  Form,
  Button,
  ComponentColor,
  ComponentStatus,
  Dropdown,
  MultiSelectDropdown,
} from 'src/reusable_ui'

// Utils
import {proxy} from 'src/utils/queryUrlGenerator'
import {parseMetaQuery} from 'src/tempVars/parsing'
import {restartable} from 'src/shared/utils/restartable'

// Types
import {RemoteDataState, Source} from 'src/types'

// This constant is selected so that dropdown menus will not overflow out of
// the `.flux-script-wizard--wizard` window
const DROPDOWN_MENU_HEIGHT = 110

interface Props {
  source: Source
  children: JSX.Element
  isWizardActive: boolean
  onSetIsWizardActive: (isWizardActive: boolean) => void
  onAddToScript: (script: string) => void
}

interface State {
  buckets: string[]
  bucketsStatus: RemoteDataState
  selectedBucket: string | null
  measurements: string[]
  measurementsStatus: RemoteDataState
  selectedMeasurement: string | null
  fields: string[]
  fieldsStatus: RemoteDataState
  selectedFields: string[] | null
}

export default class FluxScriptWizard extends PureComponent<Props, State> {
  public state = {
    buckets: [],
    bucketsStatus: RemoteDataState.NotStarted,
    selectedBucket: null,
    measurements: [],
    measurementsStatus: RemoteDataState.NotStarted,
    selectedMeasurement: null,
    fields: [],
    fieldsStatus: RemoteDataState.NotStarted,
    selectedFields: null,
  }

  private fetchBuckets = restartable(fetchBuckets)
  private fetchMeasurements = restartable(fetchMeasurements)
  private fetchFields = restartable(fetchFields)

  public componentDidMount() {
    this.fetchAndSetBuckets()
  }

  public render() {
    const {children, isWizardActive} = this.props
    const {
      buckets,
      measurements,
      fields,
      selectedBucket,
      selectedMeasurement,
      selectedFields,
    } = this.state

    if (!isWizardActive) {
      return (
        <div className="flux-script-wizard">
          <div className="flux-script-wizard--children">{children}</div>
        </div>
      )
    }

    return (
      <div className="flux-script-wizard">
        <div className="flux-script-wizard--children">{children}</div>
        <div
          className="flux-script-wizard--backdrop"
          onClick={this.handleClose}
        />
        <div className="flux-script-wizard--wizard">
          <div className="flux-script-wizard--wizard-header">
            <h3>Flux Script Wizard</h3>
            <div
              className="flux-script-wizard--close"
              onClick={this.handleClose}
            >
              &times;
            </div>
          </div>
          <div className="flux-script-wizard--wizard-body">
            <Form>
              <Form.Element label="Choose a Bucket">
                <Dropdown
                  status={this.bucketDropdownStatus}
                  selectedID={selectedBucket}
                  maxMenuHeight={DROPDOWN_MENU_HEIGHT}
                  onChange={this.handleSelectBucket}
                >
                  {buckets.map(bucket => (
                    <Dropdown.Item key={bucket} id={bucket} value={bucket}>
                      {bucket}
                    </Dropdown.Item>
                  ))}
                </Dropdown>
              </Form.Element>
              <Form.Element label="Choose a Measurement">
                <Dropdown
                  status={this.measurementDropdownStatus}
                  selectedID={selectedMeasurement}
                  maxMenuHeight={DROPDOWN_MENU_HEIGHT}
                  onChange={this.handleSelectMeasurement}
                >
                  {measurements.map(measurement => (
                    <Dropdown.Item
                      key={measurement}
                      id={measurement}
                      value={measurement}
                    >
                      {measurement}
                    </Dropdown.Item>
                  ))}
                </Dropdown>
              </Form.Element>
              <Form.Element label="Choose Measurement Fields">
                <MultiSelectDropdown
                  status={this.fieldsDropdownStatus}
                  selectedIDs={selectedFields}
                  emptyText={'All Fields'}
                  maxMenuHeight={DROPDOWN_MENU_HEIGHT}
                  onChange={this.handleSelectFields}
                >
                  {fields.map(field => (
                    <Dropdown.Item key={field} id={field} value={{id: field}}>
                      {field}
                    </Dropdown.Item>
                  ))}
                </MultiSelectDropdown>
              </Form.Element>
              <Form.Footer>
                <Button
                  text="Add to Script"
                  color={ComponentColor.Primary}
                  status={this.buttonStatus}
                  onClick={this.handleAddToScript}
                />
              </Form.Footer>
            </Form>
          </div>
        </div>
      </div>
    )
  }

  private get bucketDropdownStatus(): ComponentStatus {
    const {bucketsStatus} = this.state
    const bucketDropdownStatus = toComponentStatus(bucketsStatus)

    return bucketDropdownStatus
  }

  private get measurementDropdownStatus(): ComponentStatus {
    const {measurementsStatus} = this.state
    const measurementDropdownStatus = toComponentStatus(measurementsStatus)

    return measurementDropdownStatus
  }

  private get fieldsDropdownStatus(): ComponentStatus {
    const {fieldsStatus} = this.state
    const fieldsDropdownStatus = toComponentStatus(fieldsStatus)

    return fieldsDropdownStatus
  }

  private get buttonStatus(): ComponentStatus {
    const {bucketsStatus, measurementsStatus, fieldsStatus} = this.state
    const allDone = [bucketsStatus, measurementsStatus, fieldsStatus].every(
      s => s === RemoteDataState.Done
    )

    if (allDone) {
      return ComponentStatus.Default
    }

    return ComponentStatus.Disabled
  }

  private handleClose = () => {
    this.props.onSetIsWizardActive(false)
  }

  private handleSelectBucket = (selectedBucket: string) => {
    this.setState({selectedBucket}, this.fetchAndSetMeasurements)
  }

  private handleSelectMeasurement = (selectedMeasurement: string) => {
    this.setState({selectedMeasurement}, this.fetchAndSetFields)
  }

  private handleSelectFields = (selectedFields: string[]) => {
    this.setState({selectedFields})
  }

  private handleAddToScript = () => {
    const {onSetIsWizardActive, onAddToScript} = this.props
    const {selectedBucket, selectedMeasurement, selectedFields} = this.state
    const script = renderScript(
      selectedBucket,
      selectedMeasurement,
      selectedFields
    )

    onAddToScript(script)
    onSetIsWizardActive(false)
  }

  private fetchAndSetBuckets = async () => {
    const {source} = this.props

    this.setState({
      buckets: [],
      bucketsStatus: RemoteDataState.Loading,
      selectedBucket: null,
      measurements: [],
      measurementsStatus: RemoteDataState.NotStarted,
      selectedMeasurement: null,
      fields: [],
      fieldsStatus: RemoteDataState.NotStarted,
      selectedFields: [],
    })

    let buckets

    try {
      buckets = await this.fetchBuckets(source.links.proxy)
    } catch {
      this.setState({
        buckets: [],
        bucketsStatus: RemoteDataState.Error,
        selectedBucket: null,
      })

      return
    }

    // TODO: Replace with a user's “default database”
    const selectedBucket = buckets.includes('telegraf')
      ? 'telegraf'
      : buckets[0]

    this.setState(
      {
        buckets,
        bucketsStatus: RemoteDataState.Done,
        selectedBucket,
      },
      this.fetchAndSetMeasurements
    )
  }

  private fetchAndSetMeasurements = async () => {
    const {source} = this.props
    const {selectedBucket} = this.state

    this.setState({
      measurements: [],
      measurementsStatus: RemoteDataState.Loading,
      selectedMeasurement: null,
      fields: [],
      fieldsStatus: RemoteDataState.NotStarted,
      selectedFields: [],
    })

    let measurements

    try {
      measurements = await this.fetchMeasurements(
        source.links.proxy,
        selectedBucket
      )
    } catch {
      this.setState({
        measurements: [],
        measurementsStatus: RemoteDataState.Error,
        selectedMeasurement: null,
      })

      return
    }

    this.setState(
      {
        measurements,
        measurementsStatus: RemoteDataState.Done,
        selectedMeasurement: measurements[0],
      },
      this.fetchAndSetFields
    )
  }

  private fetchAndSetFields = async () => {
    const {source} = this.props
    const {selectedBucket, selectedMeasurement} = this.state

    this.setState({
      fields: [],
      fieldsStatus: RemoteDataState.Loading,
      selectedFields: [],
    })

    let fields

    try {
      fields = await this.fetchFields(
        source.links.proxy,
        selectedBucket,
        selectedMeasurement
      )
    } catch {
      this.setState({
        fields: [],
        fieldsStatus: RemoteDataState.Error,
        selectedFields: [],
      })

      return
    }

    this.setState({
      fields,
      fieldsStatus: RemoteDataState.Done,
      selectedFields: [],
    })
  }
}

async function fetchBuckets(proxyLink: string): Promise<string[]> {
  const query = 'SHOW DATABASES'
  const resp = await proxy({source: proxyLink, query})
  const databases = parseMetaQuery(query, resp.data)

  databases.sort()

  return databases
}

async function fetchMeasurements(
  proxyLink: string,
  bucket: string
): Promise<string[]> {
  const query = `SHOW MEASUREMENTS ON "${bucket}"`
  const resp = await proxy({source: proxyLink, query})
  const measurements = parseMetaQuery(query, resp.data)

  measurements.sort()

  return measurements
}

async function fetchFields(
  proxyLink: string,
  bucket: string,
  measurement: string
): Promise<string[]> {
  const query = `SHOW FIELD KEYS ON "${bucket}" FROM "${measurement}"`
  const resp = await proxy({source: proxyLink, query})
  const fields = parseMetaQuery(query, resp.data)

  fields.sort()

  return fields
}

function toComponentStatus(state: RemoteDataState): ComponentStatus {
  switch (state) {
    case RemoteDataState.NotStarted:
      return ComponentStatus.Disabled
    case RemoteDataState.Loading:
      return ComponentStatus.Loading
    case RemoteDataState.Error:
      return ComponentStatus.Error
    case RemoteDataState.Done:
      return ComponentStatus.Default
  }
}

function renderScript(
  selectedBucket: string,
  selectedMeasurement: string,
  selectedFields: string[]
): string {
  let filterPredicate = `r._measurement == "${selectedMeasurement}"`

  if (selectedFields.length) {
    const fieldsPredicate = selectedFields
      .map(f => `r._field == "${f}"`)
      .join(' or ')

    filterPredicate += ` and (${fieldsPredicate})`
  }

  const from = `from(bucket: "${selectedBucket}")`
  const range = `|> range(start: -1h)`
  const filter = `|> filter(fn: (r) => ${filterPredicate})`
  const script = [from, range, filter].join('\n  ')

  return script
}
