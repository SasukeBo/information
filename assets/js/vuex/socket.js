export default {
  namespaced: true,
  state: {
    deviceChannel: null,
    systemChannel: null,
  },
  mutations: {
    SET_DEVICE_CHANNEL: (state, payload) => state.deviceChannel = payload,
    SET_SYSTEM_CHANNEL: (state, payload) => state.systemChannel = payload
  },
  actions: {
    setDeviceChannel({ commit }, payload) {
      commit('SET_DEVICE_CHANNEL', payload)
    },
    setSystemChannel({ commit }, payload) {
      commit('SET_SYSTEM_CHANNEL', payload)
    }
  }
};
