interface Data {
  error: {
    status: boolean;
  };
  response: object | null;
}

class TryCatchHandler {
  private static data: Data;

  constructor() {
    TryCatchHandler.data = { error: { status: true }, response: null };
  }

  get getData(): Data {
    return this.data;
  }

  set setData(data: Data) {
    TryCatchHandler.data = data;
  }
}

export default TryCatchHandler;
