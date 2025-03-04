type validationRule = (val: any) => string | boolean;

export function required(label: string): validationRule {
  return (val: any) => {
    if (val === undefined || val === null || val === "") {
      return `${label} is required`;
    }
    return true;
  };
}
