import Axios from "axios";

const bookListApi = Axios.create({
  baseURL: "http://localhost:8080",
});

export default bookListApi;
