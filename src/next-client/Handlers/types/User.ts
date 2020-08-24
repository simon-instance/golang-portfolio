export default interface User {
    id: string;
    username: string;
    Post?: {
        title: string;
        content: string;
    };
}
