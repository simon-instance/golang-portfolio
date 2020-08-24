import * as TryCatch from "./Errors/TryCatchHandler";
import * as Request from "./RequestHandler";

// TryCatch
export type TryCatchProps = TryCatch.HandlerProps;
export type TryCatchDataProps = TryCatch.DataProps;
export const TryCatchHandler = TryCatch.Handler;

// Request
export type RequestProps = Request.HandlerProps;
export const RequestHandler = Request.Handler;
