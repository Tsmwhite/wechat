export const SetsessionStorage = (key, value) => {
  window.sessionStorage.setItem(key, value);
};

export const GetsessionStorage = (key) => {
  return window.sessionStorage.getItem(key);
};

export const SetlocalStorage = (key, value) => {
  window.localStorage.setItem(key, value);
};

export const GetlocalStorage = (key) => {
  return window.localStorage.getItem(key);
};

export const RemovelocalStorage = (key) => {
  return window.localStorage.removeItem(key);
};
