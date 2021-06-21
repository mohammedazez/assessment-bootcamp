import { combineReducers } from "redux";

import userRegisterReducer from "../redux/user/register/userRegisterReducer";
import listReducers from "../redux/lists/listsReducers";

const rootReducer = combineReducers({
  userRegister: userRegisterReducer,
  listReducer: listsReducers,
});

export default rootReducer;
