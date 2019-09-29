export default {
  namespaced: true,
  state: {
    status: null,
    avatarURL: null,
    phone: null,
    role: null,
    name: null,
    email: null
  },
  mutations: {
    SET_AVATARURL: (state, payload) => state.avatarURL = payload,
    SET_PHONE: (state, payload) => state.phone = payload,
    SET_ROLE: (state, payload) => state.role = payload,
    SET_NAME: (state, payload) => state.name = payload,
    SET_EMAIL: (state, payload) => state.email = payload,
    SET_STATUS: (state, payload) => state.status = payload,
    LOGOUT: (state, callback) => {
      state.uuid = null;
      callback()
    }
  },
  actions: {
    setUserData({ commit }, payload) {
      commit('SET_AVATARURL', payload.avatarURL)
      commit('SET_PHONE', payload.phone)
      commit('SET_ROLE', payload.role)
      commit('SET_NAME', payload.name)
      commit('SET_EMAIL', payload.email)
      commit('SET_STATUS', payload.status)
    },
    clearUserData({ commit }) {
      commit('SET_AVATARURL', null)
      commit('SET_PHONE', null)
      commit('SET_ROLE', null)
      commit('SET_NAME', null)
      commit('SET_EMAIL', null)
      commit('SET_STATUS', null)
    },
    logout({ commit }, callback) {
      commit('LOGOUT', callback)
    }
  }
};
