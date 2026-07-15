declare module '@apiverve/qrcodegenerator' {
  export interface qrcodegeneratorOptions {
    api_key: string;
    secure?: boolean;
  }

  /**
   * Describes fields the current plan does not unlock. Locked fields arrive as null
   * in `data`; `locked_fields` names them, using dot paths for nested fields.
   * Absent when the plan unlocks everything.
   */
  export interface PremiumInfo {
    message: string;
    upgrade_url: string;
    locked_fields: string[];
  }

  export interface qrcodegeneratorResponse {
    status: string;
    error: string | null;
    data: QRCodeGeneratorData;
    code?: number;
    premium?: PremiumInfo;
  }


  interface QRCodeGeneratorData {
      id:          null | string;
      format:      null | string;
      type:        null | string;
      correction:  null | string;
      size:        number | null;
      margin:      number | null;
      expires:     number | null;
      downloadURL: null | string;
  }

  export default class qrcodegeneratorWrapper {
    constructor(options: qrcodegeneratorOptions);

    execute(callback: (error: any, data: qrcodegeneratorResponse | null) => void): Promise<qrcodegeneratorResponse>;
    execute(query: Record<string, any>, callback: (error: any, data: qrcodegeneratorResponse | null) => void): Promise<qrcodegeneratorResponse>;
    execute(query?: Record<string, any>): Promise<qrcodegeneratorResponse>;
  }
}
