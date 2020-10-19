FROM scratch
COPY ./main /main
COPY ./conf.json /conf.json
CMD ["/main"]