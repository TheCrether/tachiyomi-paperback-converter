declare interface ResultWithNoError<T> {
	value: T;
}

declare interface ResultWithError<T> extends ResultWithNoError<T> {
	errorMessage?: string;
	error?: Error;
}

declare type Result<T> = ResultWithNoError<T> | ResultWithError<T>;

declare function convertTachiyomi(bytes: Uint8Array): Result<string>;
declare function convertPaperback(json: string): Result<Uint8Array>;