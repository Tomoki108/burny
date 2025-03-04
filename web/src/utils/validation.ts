type validationRule = (val: any) => string | boolean;

export function newRule(label: string) {
  return new Rule(label);
}

export class Rule {
  label: string;
  rules: validationRule[] = [];

  constructor(label: string) {
    this.label = label;
  }

  required(): this {
    this.rules.push((val: any) => {
      if (val === undefined || val === null || val === "") {
        return `${this.label} is required`;
      }
      return true;
    });
    return this;
  }

  oneOf(options: any[]): this {
    this.rules.push((val: any) => {
      if (!options.includes(val)) {
        return `${this.label} must be one of ${options.join(", ")}`;
      }
      return true;
    });
    return this;
  }

  lte(max: number): this {
    this.rules.push((val: any) => {
      switch (typeof val) {
        case "string":
          val as string;
          if (val.length > max) {
            return `${this.label} must be less than or equal to ${max} characters`;
          }
          return true;
        case "number":
          val as number;
          if (val > max) {
            return `${this.label} must be less than or equal to ${max}`;
          }
          return true;
        default:
          throw new Error("Invalid type");
      }
    });

    return this;
  }

  gt(min: number): this {
    this.rules.push((val: any) => {
      switch (typeof val) {
        case "string":
          val as string;
          if (val.length <= min) {
            return `${this.label} must be greater than ${min} characters`;
          }
          return true;
        case "number":
          val as number;
          if (val <= min) {
            return `${this.label} must be greater than ${min}`;
          }
          return true;
        default:
          throw new Error("Invalid type");
      }
    });
    return this;
  }
}

export interface vForm {
  validate(): Promise<vFormValidateResult>;
}

export interface vFormValidateResult {
  valid: boolean;
  errors: string[];
}
