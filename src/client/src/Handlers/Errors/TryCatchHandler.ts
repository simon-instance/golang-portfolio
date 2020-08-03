interface User {
  id: string;
  username: string;
  Post?: {
    title: string;
    content: string;
  };
}

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

export interface DataInterface {
  error: {
    status: boolean;
    message: string | undefined;
  };
  response: Response;
}

export interface Interface {
  Data: DataInterface;

  handleData(data: DataInterface): void;
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

  public handleData(data: DataInterface): void {
    console.log(data);
  }
}
