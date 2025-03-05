import { reactive, ref } from "vue";

export type AlertType = "info" | "success" | "error" | "warning" | undefined;

export interface AlertContext {
  show: boolean;
  text: string;
  type: AlertType;
}

const alertCtx = reactive<AlertContext>({
  show: false,
  text: "",
  type: undefined,
});

const alertShow = ref(false);
const alertText = ref("");
const alertType = ref<AlertType>("info");

export const useAlertComposable = () => {
  const alert = (msg: string, type: AlertType) => {
    alertCtx.show = true;
    alertCtx.text = msg;
    alertCtx.type = type;

    if (type === "success") {
      setTimeout(() => {
        alertCtx.show = false;
        alertCtx.text = "";
        alertCtx.type = undefined;
      }, 2000);
    }
  };

  return {
    alertCtx,
    alert,
  };
};
