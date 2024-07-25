import { CONFIG_FILE_LOADED } from "@/constants/event-names/CONFIG_FILE_LOADED"

class ConfigFileLoadedEvent extends EventTarget {
  constructor() {
    super()
  }

  private _configFileLoaded = new Event(CONFIG_FILE_LOADED)

  public configFileLoaded() {
    this.dispatchEvent(this._configFileLoaded)
  }
}

export const configFileLoaded = new ConfigFileLoadedEvent()