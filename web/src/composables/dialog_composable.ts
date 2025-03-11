import { reactive } from "vue";

export interface DialogContext {
  show: boolean;
  title: string;
  text: string;
}

const dialogCtx = reactive<DialogContext>({
  show: false,
  title: "",
  text: "",
});

export const useDialogComposable = () => {
  const dialog = (title: string, msg: string) => {
    dialogCtx.show = true;
    dialogCtx.title = title;
    dialogCtx.text = msg;
  };
  return {
    dialogCtx,
    dialog,
  };
};
