const initialState = {
  lists: [],
  list: null,
  isLoading: false,
  error: null,
};

const listsReducers = (state = initialState, action) => {
  switch (action.type) {
    case "LIST_LOADING":
      return {
        ...state,
        isLoading: true,
      };
    case "FETCH_lIST":
      return {
        ...state,
        isLoading: false,
        lists: action.payload,
      };
    case "CREATE_lIST":
      return {
        ...state,
        isLoading: false,
        lists: action.payload,
      };
    case "UPDATE_lIST":
      return {
        ...state,
        isLoading: false,
        lists: action.payload,
      };
    case "DELETE_lIST":
      return {
        ...state,
        isLoading: false,
        list: null,
      };
    case "ERROR_lIST":
      return {
        ...state,
        isLoading: false,
      };
    default:
      break;
  }
};

export default listsReducers;
