export const API_BASE_URL =
  import.meta.env?.VITE_API_BASE_URL || "http://localhost:1323/api/v1";

export function getAuthHeader(): HeadersInit {
  const token = localStorage.getItem("token");
  if (!token) {
    throw new Error("No token found");
  }

  return {
    "Content-Type": "application/json",
    Authorization: `Bearer ${token}`,
  };
}

// yyyy-mm-ddを、ISO 8601形式（例: "2025-02-14T00:00:00.000Z"）に変換する replacer
// リクエストに使う日付はISO 8601形式でないとAPI側がパースできない
export function replaceDateWithISOString(key: string, value: any): any {
  if (key === "start_date") {
    return new Date(value).toISOString();
  }
  return value;
}

export class ErrorResponse {
  message: string = "";
  details?: {
    field: string;
    message: string;
  }[];

  getMessage(): string {
    // messageとdetailsを結合して返す
    if (this.details) {
      return `${this.message}: ${this.details
        .map((d) => `${d.field} ${d.message}`)
        .join(", ")}`;
    }
    return this.message;
  }
}
