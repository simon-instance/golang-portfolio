import AuthFields from "./AuthFields";
import User from "./User";

//
// When api returns 404, data.message = api error
// When api returns a 200 status, data.User = api User (from firebase database) (the user is found and authenticated)
//
export default interface Response {
    data:
        | {
              fields?: Fields;
              message?: string;
              User?: User;
          }
        | undefined;
    message: string | undefined;
    status: number | undefined;
}
