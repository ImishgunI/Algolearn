FROM nginx:alpine

RUN rm -rf /usr/share/nginx/html/*

COPY ./web/html /usr/share/nginx/html/
COPY ./web/css /usr/share/nginx/html/css/
COPY ./web/js /usr/share/nginx/html/js/

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]
