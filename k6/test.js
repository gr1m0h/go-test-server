import http from "k6/http";

export default () => {
    http.get("http://localhost:8080/tests");
};
