import Axios from "axios";

const webListApi = Axios.create({
  baseURL: "http://localhost:8080",
});

export default webListApi;
