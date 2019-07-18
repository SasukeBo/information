function load (path) {
  const component = () => import(`js/vue/${path}`);
  return component;
};

export {
  load
}
