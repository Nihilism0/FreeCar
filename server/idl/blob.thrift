namespace go blob

struct CreateBlobRequest {
    1: i64 account_id
    2: i32 upload_url_timeout_sec
}

struct CreateBlobResponse {
    1: i64 id
    2: string upload_url
}

struct GetBlobURLRequest {
    1: i64 id
    2: i32 timeout_sec
}

struct GetBlobURLResponse {
    1: string url
}

service BlobService {
    CreateBlobResponse CreateBlob (1:CreateBlobRequest req)
    GetBlobURLResponse GetBlobURL(1: GetBlobURLRequest req)
}
