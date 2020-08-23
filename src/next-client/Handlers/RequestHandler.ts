import { TryCatchHandler, TryCatchInterface, TryCatchDataInterface } from "./";

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

    makeRequest(): any {
        const TCHandler: TryCatchInterface = new TryCatchHandler();

        try {
            const response: any = await fetch(
                `${Handler.baseUrl}${
                    this.uri.endsWith("/") ? this.uri : this.uri + "/"
                }`,
                {
                    method: this.method,
                    body: JSON.stringify(this.data),
                }
            );

            let data: TryCatchDataInterface = TCHandler.Data;

            //TODO

            data.response.data = await response.json();
            data.response.status = response.status;

            const error: Error | void = TCHandler.handleData(data);
            if (typeof error !== "undefined") throw error;
        } catch (e) {
            data.error.status = true;
            data.error.message = e;
        } finally {
            console.log(data.error.status);
            TCHandler.Data = data;
        }
    }
}
