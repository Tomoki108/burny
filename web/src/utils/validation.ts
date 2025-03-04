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
      if (val > max) {
        return `${this.label} must be less than or equal to ${max}`;
      }
      return true;
    });
    return this;
  }

  gt(min: number): this {
    this.rules.push((val: any) => {
      if (val <= min) {
        return `${this.label} must be greater than ${min}`;
      }
      return true;
    });
    return this;
  }
}

export interface vForm {
  validate(): vFormValidateResult;
}

export interface vFormValidateResult {
  valid: boolean;
  errors: string[];
}
