import { ref } from "vue";

export type AlertType = "info" | "success" | "error" | "warning";

const alertShow = ref(false);
const alertText = ref("");
const alertType = ref<AlertType>("info");

export const useAlertComposable = () => {
  const alert = (msg: string, type: AlertType) => {
    alertText.value = msg;
    alertType.value = type;
    alertShow.value = true;

    // setTimeout(() => {
    //   alertShow.value = false;
    //   alertText.value = "";
    //   alertType.value = "info";
    // }, 5000);
  };

  return {
    alertShow,
    alertText,
    alertType,
    alert,
  };
};
