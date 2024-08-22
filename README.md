# Cloud Run `gRPC Proxy` sidecar for Google APIs

---

## Motivation:

During development, it is often useful to perform Google APIs traffic analysis at the protocol layer in order to troubleshoot specific/gnarly application related conditions/issues.

For `gRPC` based API clients, getting information about what's actually being sent and received and, how networking and latency are behaving,
it can be not as simple to get all the datapoints required to troubleshoot without performing complex client configurations and instrumentation.

The `gRPC Proxy` for Google APIs aims to simplify the troubleshooting process by providing all the information required by developers leveraging `gRPC` clients to perform Google APIs operations;
for example: if there is a timeout, was the actual `rpc` sent or not?; if there was an unexpected response: was the actual request correct?, etc...

![alt text](https://github.com/gchux/cloud-run-grpc-proxy/blob/main/img/grpc_proxy.png?raw=true)

---

## Features

- Intercept ALL gRPC traffic regardless of the `rpc` and `message`.

- Show all information from inbound at outbound `rpc` metadata.

- Show all information from IP and TCP layers on the inbound and outbound directions.

- Capture stream setup latency, along side latency for every `rpc` performed within the stream.

- Capture and log `rpc` request and response messages ( only if the [plugin](https://github.com/gchux/cloud-run-grpc-proxy/tree/main/_plugins) for the package is available; otherwise, it will log the [`Empty message`](https://pkg.go.dev/google.golang.org/protobuf/types/known/emptypb).

- ALL `rpc`s within the same stream are tagged with the same serial number to easily correlate; additionally, IDs are generated for the client socket, server socket, and proxy flow.

- Log events into Cloud Logging and optionally add trace and span id to perform distribyted trace analysis in Cloud Trace.

- Handler authentication for all `rpc`s executed through the `gRPC Proxy` server, unless the `rpc` already contains an `Authorization` header.

---

## Available `configurations`

The `gRPC Proxy` sidecar accepts the following environment variables:

- `PROJECT_ID`: (STRING, _optional_) the Google Cloud Platform project ID that will be used for header [`x-goog-user-project`](https://cloud.google.com/apis/docs/system-parameters#definitions).

- `GRPC_PROXY_PORT`: (NUMBER, _optional_) the TCP port number at which the gRPC Proxy server will listen for connections; default value is `5001`.

- `GRPC_PROXY_TARGET_HOST`: (STRING, _optional_) the fully qualified domain name of the upstream/target host; i/e: `run.googleapis.com`.

- `GRPC_PROXY_TARGET_PORT`: (NUMBER, _optional_) the TCP port at which the upstream/target API server accepts connections; i/e: `443`.

Additionally, the `gRPC Proxy` accepts runtime configurations in the form of the following headers:

- `x-grpc-proxy-endpoint`: (STRING, _optional_) it must be the concatenation of `GRPC_PROXY_TARGET_HOST` and `GRPC_PROXY_TARGET_PORT` separated by a colon; i/e: `run.googleapis.com:443`.

- `x-grpc-proxy-project`: (STRING, _optional_) same as the environment variable `PROJECT_ID`.

- `x-grpc-proxy-location`: (STRING, _optional_) the GCP location for the intended resource being affected by the API operation.

- `x-cloud-trace-context`: (STRING, _optional_) the trace and span IDs to be used when generating logs for Cloud Logging; it allows for Cloud Trace integration.

**NOTE**: if both, the environment variable and the header are available, the header always takes precedence ( it overrides the environment variable ).

---

## How it works:

The sidecar uses:

- [`google-cloud-go`](https://github.com/googleapis/google-cloud-go) compiled PROTOs from most client libraries.

- [`googleapis`](https://github.com/googleapis/googleapis) PROTO definitions to generate plugins code for most Google APIS.

- [`gRPC GO`](https://pkg.go.dev/google.golang.org/grpc) implementation to create the `gRPC Server` and register the [`UnknownServiceHandler`](https://pkg.go.dev/google.golang.org/grpc#UnknownServiceHandler) to catch ALL `rpc`s.

---

## How to use:

- Pull the latest version of the pre-built sidecar container image: `docker pull ghcr.io/gchux/cloud-run-grpc-proxy:latest`

- Tag the image with the expected Artifact Registry repository URI, i/e: `docker tag ghcr.io/gchux/cloud-run-grpc-proxy:latest us-docker.pkg.dev/${PROJECT_ID}/${GCP_AR_REPO}/grpc-proxy:latest`

- Push the image into the expected Artifact Registry repository; i/e: `docker push us-docker.pkg.dev/${PROJECT_ID}/${GCP_AR_REPO}/grpc-proxy:latest`

- Deploy a new revision of the Cloud Run service using the `gRPC Proxy` sidecar; i/e:

      ```sh
      gcloud beta run deploy ${SERVICE_NAME} \
        --project=${PROJECT_ID} \
        --region=${SERVICE_REGION} \
        --service-account=${SERVICE_ACCOUNT} \
        --container=${INGRESS_CONTAINER_NAME}-1 \
        --image=${INGRESS_IMAGE_URI} \
        --port=${INGRESS_PORT} \
        --container=${GRPC_PROXY_SIDECAR_NAME}-1 \
        --image=${GRPC_PROXY_IMAGE_URI} \
        --cpu=1 --memory=1G \
        --set-env-vars="PROJECT_ID=${PROJECT_ID},GRPC_PROXY_PORT=${GRPC_PROXY_PORT},GRPC_PROXY_TARGET_HOST=${GRPC_PROXY_TARGET_HOST},GRPC_PROXY_TARGET_PORT=${GRPC_PROXY_TARGET_PORT}" \
      ```

---

## How to use `plugins`

In order to keep the container image simple, `proto` definitions for most Google APIs are bundled as plugins which are laoded by the `gRPC Proxy` sidecar at runtime.

Plugins are build using [`Go Plugins package`](https://pkg.go.dev/plugin) and, `gRPC Proxy` plugins are defined using [`pkg` files](https://github.com/gchux/cloud-run-grpc-proxy/tree/main/_plugins_pkg) which contains all the `PROTO packages` that contain the intended to be used services.

This `package file` strategy allows developers to create a bundle with multiple packages from different APIs and point all the clients to the same `gRPC Proxy` server at runtime.

The image provided does not contain any bundled plugins, it is a base image in which the `gRPC Proxy` handles all `rpc`s and all logging entries contain the [`Epmty message`](https://pkg.go.dev/google.golang.org/protobuf/types/known/emptypb); however, all other relevant information about the `rpc` will be available.
