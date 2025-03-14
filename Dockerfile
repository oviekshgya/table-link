FROM build-base AS builder

ARG GITLAB_USERNAME
ARG PWDG

WORKDIR /app

RUN git config --global url."https://${GITLAB_USERNAME}:${PWDG}@gitlab.com".insteadOf "https://gitlab.com"

COPY go.mod go.sum ./

# Optional: Uncomment if necessary
# RUN go mod tidy

COPY . .

RUN go build -o table-link

FROM alpine:3.17 AS runner

WORKDIR /app

COPY --from=builder /app/table-link .

EXPOSE ${SERVICE_PORT}

CMD ["/app/table-link"]

