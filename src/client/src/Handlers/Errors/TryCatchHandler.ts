interface User {
  id: string;
  username: string;
  Post?: {
    title: string;
    content: string;
  };
}

//
// When api returns 404, data.message = api error
// When api returns a 200 status, data.User = api User (from firebase database) (the user is found and authenticated)
//
interface Response {
  data:
    | {
        message?: string;
        User?: User;
      }
    | undefined;
  message: string | undefined;
  status: number | undefined;
}

// message = readable user error
export interface DataInterface {
  error: {
    status: boolean;
    message: string | undefined;
  };
  response: Response;
}

export interface Interface {
  Data: DataInterface;

  handleData(data: DataInterface): void | Error;
}

export class Handler implements Interface {
  public Data: DataInterface = {
    error: { status: false, message: undefined },
    response: {
      data: undefined,
      message: undefined,
      status: undefined,
    },
  };

  public handleData(data: DataInterface): void | Error {
    this.Data = data;

    if (
      data.response.data !== undefined &&
      data.response.status !== undefined
    ) {
      if (data.response.data.message !== undefined) {
        this.Data.response.message = data.response.data.message;

        if (data.response.status === 404 || data.response.status === 401)
          return new Error(this.Data.response.message);

        return new Error("Geen omschrijving voor deze foutmelding");
      }
      if (data.response.status === 200) return;
    }

    return new Error("Geen connectie met de database");
  }
}
