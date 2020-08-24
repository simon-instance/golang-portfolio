import ApiResponse from "../types/ApiResponse";

// message = readable user error
export interface DataProps {
    error: {
        status: boolean;
        message: string | undefined;
    };
    response: ApiResponse;
}

export interface HandlerProps {
    Data: DataProps;

    handleData(data: DataProps): void | Error;
}

export class Handler implements HandlerProps {
    public Data: DataProps = {
        error: { status: false, message: undefined },
        response: {
            data: undefined,
            message: undefined,
            status: undefined,
        },
    };

    public handleData(data: DataProps): void | Error {
        this.Data = data;

        if (
            data.response.data !== undefined &&
            data.response.status !== undefined
        ) {
            if (data.response.data.message !== undefined) {
                this.Data.response.message = data.response.data.message;

                if (
                    data.response.status === 404 ||
                    data.response.status === 401
                )
                    return new Error(this.Data.response.message);

                return new Error("Geen omschrijving voor deze foutmelding");
            }
            if (data.response.status === 200) return;
        }

        return new Error("Geen connectie met de database");
    }
}
