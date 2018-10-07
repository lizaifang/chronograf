// Libraries
import React, {PureComponent} from 'react'

// Components
import SchemaExplorer from 'src/flux/components/SchemaExplorer'
import TimeMachineEditor from 'src/flux/components/TimeMachineEditor'
import FluxScriptWizard from 'src/shared/components/TimeMachine/FluxScriptWizard'
import Threesizer from 'src/shared/components/threesizer/Threesizer'
import {
  Button,
  ComponentSize,
  ComponentColor,
  ComponentStatus,
} from 'src/reusable_ui'

// Constants
import {HANDLE_VERTICAL} from 'src/shared/constants'

// Utils
import {getSuggestions, getAST} from 'src/flux/apis'
import {restartable} from 'src/shared/utils/Restartable'
import DefaultDebouncer, {Debouncer} from 'src/shared/utils/debouncer'
import {parseError} from 'src/flux/helpers/scriptBuilder'

// Types
import {NotificationAction, Source} from 'src/types'
import {Suggestion, Links, ScriptStatus} from 'src/types/flux'

const CHECK_SCRIPT_DELAY = 600
const VALID_SCRIPT_STATUS = {type: 'success', text: ''}

interface Props {
  script: string
  onChangeScript: (script: string) => void
  source: Source
  onUpdateStatus?: (status: ScriptStatus) => void
  links: Links
  notify: NotificationAction
}

interface State {
  suggestions: Suggestion[]
  draftScript: string
  draftScriptStatus: ScriptStatus
  hasChangedScript: boolean
  isWizardActive: boolean
}

class FluxQueryMaker extends PureComponent<Props, State> {
  private debouncer: Debouncer = new DefaultDebouncer()
  private getAST = restartable(getAST)

  public constructor(props: Props) {
    super(props)

    this.state = {
      suggestions: [],
      draftScript: props.script,
      draftScriptStatus: {type: 'none', text: ''},
      hasChangedScript: false,
      isWizardActive: false,
    }
  }

  public componentDidMount() {
    this.fetchSuggestions()
    this.checkDraftScript()
  }

  public render() {
    const {notify, source} = this.props
    const {
      suggestions,
      draftScript,
      draftScriptStatus,
      hasChangedScript,
      isWizardActive,
    } = this.state

    const submitStatus = hasChangedScript
      ? ComponentStatus.Default
      : ComponentStatus.Disabled

    const divisions = [
      {
        name: 'Script',
        size: 0.66,
        headerOrientation: HANDLE_VERTICAL,
        headerButtons: [
          <Button
            key={0}
            text={'Script Wizard'}
            onClick={this.handleShowWizard}
            color={ComponentColor.Primary}
            size={ComponentSize.ExtraSmall}
          />,
          <Button
            key={1}
            text={'Run Script'}
            onClick={this.handleSubmitScript}
            size={ComponentSize.ExtraSmall}
            status={submitStatus}
          />,
        ],
        menuOptions: [],
        render: visibility => (
          <TimeMachineEditor
            status={draftScriptStatus}
            script={draftScript}
            visibility={visibility}
            suggestions={suggestions}
            onChangeScript={this.handleChangeDraftScript}
          />
        ),
      },
      {
        name: 'Explore',
        size: 0.34,
        headerButtons: [],
        menuOptions: [],
        render: () => <SchemaExplorer source={source} notify={notify} />,
        headerOrientation: HANDLE_VERTICAL,
      },
    ]

    return (
      <FluxScriptWizard
        source={source}
        isWizardActive={isWizardActive}
        onSetIsWizardActive={this.handleSetIsWizardActive}
        onAddToScript={this.handleAddToScript}
      >
        <Threesizer
          orientation={HANDLE_VERTICAL}
          divisions={divisions}
          containerClass="page-contents"
        />
      </FluxScriptWizard>
    )
  }

  private handleSubmitScript = () => {
    const {onChangeScript, onUpdateStatus} = this.props
    const {draftScript, draftScriptStatus} = this.state

    onChangeScript(draftScript)

    if (onUpdateStatus) {
      onUpdateStatus(draftScriptStatus)
    }

    this.setState({hasChangedScript: false})
  }

  private handleShowWizard = (): void => {
    this.setState({isWizardActive: true})
  }

  private handleSetIsWizardActive = (isWizardActive: boolean): void => {
    this.setState({isWizardActive})
  }

  private handleAddToScript = (newScript: string): void => {
    const {draftScript} = this.state

    let newDraftScript = newScript

    if (draftScript.trim() !== '') {
      newDraftScript += `\n\n${draftScript}`
    }

    this.handleChangeDraftScript(newDraftScript)
  }

  private handleChangeDraftScript = async (
    draftScript: string
  ): Promise<void> => {
    this.setState(
      {
        draftScript,
        hasChangedScript: true,
      },
      () => this.debouncer.call(this.checkDraftScript, CHECK_SCRIPT_DELAY)
    )
  }

  private checkDraftScript = async () => {
    const {draftScript} = this.state

    if (draftScript.trim() === '') {
      // Don't attempt to validate an empty script
      this.setState({draftScriptStatus: VALID_SCRIPT_STATUS})

      return
    }

    let draftScriptStatus: ScriptStatus

    try {
      await this.getAST({url: this.props.links.ast, body: draftScript})

      draftScriptStatus = VALID_SCRIPT_STATUS
    } catch (error) {
      draftScriptStatus = parseError(error)
    }

    this.setState({draftScriptStatus})
  }

  private fetchSuggestions = async (): Promise<void> => {
    const {links} = this.props
    const suggestions = await getSuggestions(links.suggestions)

    this.setState({suggestions})
  }
}

export default FluxQueryMaker
