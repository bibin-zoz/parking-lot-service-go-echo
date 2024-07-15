FROM golang:1.22-alpine AS FirstBuild
WORKDIR /parking-lot-service/
COPY ./ ./
RUN go mod download
RUN go build -o ./cmd/build ./cmd/main.go


FROM scratch
COPY --from=FirstBuild /parking-lot-service/cmd/build /
COPY --from=FirstBuild /parking-lot-service/.env /
EXPOSE 8080
CMD [ "/build" ]