import { SIDE_PANEL_TOGGLED } from "@/constants/event-names/SIDE_PANEL_TOGGLED"

type SidePanelState = 'expanded' | 'collapsed'

class ToggleSidePanelEvent extends EventTarget {
  constructor() {
    super()
  }

  private _sidePanelToggled = new Event(SIDE_PANEL_TOGGLED)

  public sidePanelToggled() {
    this.dispatchEvent(this._sidePanelToggled)
  }
}

export const sidePanelToggled = new ToggleSidePanelEvent()