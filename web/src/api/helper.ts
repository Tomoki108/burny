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
