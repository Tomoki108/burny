import { reactive } from "vue";

export interface ApikeyContext {
  show: boolean;
  rawKey: string;
}

const apikeyCtx = reactive<ApikeyContext>({
  show: false,
  rawKey: "",
});

export const useApikeyComposable = () => {
  const showRawKey = (key: string) => {
    apikeyCtx.show = true;
    apikeyCtx.rawKey = key;
  };
  return {
    apikeyCtx,
    showRawKey,
  };
};
