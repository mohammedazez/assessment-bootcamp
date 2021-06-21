import bookListApi from "../../../APIs/bookListApi";

export const registerUser = (payload) => {
  return async (dispatch) => {
    try {
      dispatch({
        type: "USER_LOADING",
      });

      const response = await bookListApi({
        method: "POST",
        url: "/users/register",
        data: payload,
      });

      console.log(response);

      return dispatch({
        type: "USER_REGISTER",
        payload: response,
      });
    } catch (error) {
      console.log(error);
    }
  };
};

export const loginUser = (payload) => {
  return async (dispatch) => {
    try {
      dispatch({
        type: "USER_LOADING",
      });

      const response = await bookListApi({
        method: "POST",
        url: "/users/login",
        data: payload,
      });

      console.log(response);

      return dispatch({
        type: "USER_LOGIN",
        payload: response,
      });
    } catch (error) {
      console.log(error);
    }
  };
};
