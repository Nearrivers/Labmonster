import { SETTINGS_OPENED } from "@/constants/event-names/SETTINGS_OPENED"

class OpenSettingsEvent extends EventTarget {

  constructor() {
    super()
  }

  private _settingsOpened = new Event(SETTINGS_OPENED)

  public openSettings() {
    this.dispatchEvent(this._settingsOpened)
  }
}

export const settingsOpened = new OpenSettingsEvent()