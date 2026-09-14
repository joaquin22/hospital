import { configureStore } from "@reduxjs/toolkit";
import reducer from "./reducer"; // Ensure this is correctly imported

const store = configureStore({
  reducer, // Correctly passing the reducer
  middleware: (getDefaultMiddleware) => getDefaultMiddleware() // Ensures compatibility
});

export default store;
