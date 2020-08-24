import { UserFields } from "./FormFields";

export default interface User {
    id: string;
    username: string;
    Post?: {
        title: string;
        content: string;
    };
    fields?: UserFields;
}
