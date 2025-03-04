import { defineStore } from "pinia";

interface Alert {
  id: number;
  message: string;
  timeout: number;
  props: Record<string, any>;
}

export const useAlertStore = defineStore("alert", {
  state: () => ({
    alerts: [] as Alert[], // アラートのリスト
  }),
  actions: {
    // アラートを追加する
    addAlert(alert: Omit<Alert, "id">) {
      const id = Date.now();
      const newAlert = { ...alert, id };
      this.alerts.push(newAlert);
      if (alert.timeout) {
        setTimeout(() => {
          this.removeAlert(id);
        }, alert.timeout);
      }
    },
    // 特定の ID のアラートを削除する
    removeAlert(id: number) {
      this.alerts = this.alerts.filter((alert) => alert.id !== id);
      if (this.alerts.length === 0) {
        this.$reset();
      }
    },
    // すべてのアラートをクリアする
    clearAlerts() {
      this.alerts = [];
    },
  },
});
