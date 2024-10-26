import GameSettings from "@/components/settings/GameSettings.vue";
import GeneralSettings from "@/components/settings/GeneralSettings.vue";
import GraphSettings from "@/components/settings/GraphSettings.vue";
import MediaSettings from "@/components/settings/MediaSettings.vue";
import { SETTINGS_OPENED } from "@/constants/event-names/SETTINGS_OPENED";
import { settingsOpened } from "@/events/OpenSettings";
import { DefinedComponent } from "@vue/test-utils/dist/types";
import { useMagicKeys } from "@vueuse/core";
import { Wrench, Gamepad2, MonitorPlay, GitFork } from "lucide-vue-next";
import { watch, FunctionalComponent, ref } from "vue";
import { useEventListener } from "../useEventListener";

type SettingsTab = {
  icon: FunctionalComponent;
  tab: DefinedComponent;
};

export function useSettingsDialog(openDialog: () => void) {
  const isDialogOpen = ref(false);
  const currentTab = ref('Jeux');
  const tabs = new Map<string, SettingsTab>();
  tabs.set(GeneralSettings.name!, { icon: Wrench, tab: GeneralSettings });
  tabs.set(GameSettings.name!, { icon: Gamepad2, tab: GameSettings });
  tabs.set(MediaSettings.name!, { icon: MonitorPlay, tab: MediaSettings });
  tabs.set(GraphSettings.name!, { icon: GitFork, tab: GraphSettings });

  const keys = useMagicKeys();
  const CtrlComma = keys['Ctrl+,'];
  useEventListener(settingsOpened, SETTINGS_OPENED, openDialog);

  watch(CtrlComma, async (v) => {
    if (!v) {
      return;
    }
    isDialogOpen.value = true;
  });

  return {
    isDialogOpen,
    currentTab,
    tabs
  }
}