FROM ubuntu:20.04
COPY ./dist/url_shortener_api /url_shortener_api
RUN chmod +x /url_shortener_api
CMD /url_shortener_api