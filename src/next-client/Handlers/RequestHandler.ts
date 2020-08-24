import { TryCatchHandler, TryCatchProps, TryCatchDataProps } from "./";

export interface HandlerProps {
    data: object | undefined;
    method?: "POST" | "PUT" | "GET" | "DELETE";
    uri: string;

    makeRequest: () => any;
}

export class Handler implements HandlerProps {
    data;
    method;
    uri;
    private static baseUrl = process.env.NODE_REQUEST_CONF;

    RequestHandler() {
        this.method = "GET";
        this.data = undefined;
    }

    async makeRequest(): Promise<TryCatchDataProps> {
        const TCHandler: TryCatchProps = new TryCatchHandler();

        let data: TryCatchDataProps = TCHandler.Data;
        try {
            const response: any = await fetch(`${Handler.baseUrl}${this.uri}`, {
                method: this.method,
                body: JSON.stringify(this.data),
            });

            data.response.data = await response.json();
            data.response.status = response.status;

            const error: Error | void = TCHandler.handleData(data);
            if (typeof error !== "undefined") throw error;
        } catch (e) {
            data.error.status = true;
            data.error.message = e;
        }
        return data;
    }
}
