import { reactive } from "vue";

export interface DialogContext {
  show: boolean;
  text: string;
}

const dialogCtx = reactive<DialogContext>({
  show: false,
  text: "",
});

export const useDialogComposable = () => {
  const dialog = (msg: string) => {
    dialogCtx.show = true;
    dialogCtx.text = msg;
  };
  return {
    dialogCtx,
    dialog,
  };
};
