
class OpenSettingsEvent extends EventTarget {

  constructor() {
    super()
  }

  private _settingsOpened = new Event('settingsOpened')

  public openSettings() {
    this.dispatchEvent(this._settingsOpened)
  }
}

export const openSettingEvent = new OpenSettingsEvent()