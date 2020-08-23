import * as TryCatch from "./Errors/TryCatchHandler";
import * as Request from "./RequestHandler";

// TryCatch
export type TryCatchInterface = TryCatch.Interface;
export type TryCatchDataInterface = TryCatch.DataInterface;
export const TryCatchHandler = TryCatch.Handler;

// Request
export type RequestHandlerProps = Request.HandlerProps;
export const RequestHandler = Request.Handler;
