export default {
  namespaced: true,
  state: {
    status: null,
    uuid: null,
    avatarURL: null,
    phone: null,
    role: null,
    userExtend: null,
  },
  mutations: {
    SET_USERUUID: (state, payload) => state.uuid = payload,
    SET_AVATARURL: (state, payload) => state.avatarURL = payload,
    SET_PHONE: (state, payload) => state.phone = payload,
    SET_ROLE: (state, payload) => state.role = payload,
    SET_USER_EXTEND: (state, payload) => state.userExtend = payload,
    SET_STATUS: (state, payload) => state.status = payload,
    LOGOUT: (state, callback) => {
      state.uuid = null;
      callback()
    }
  },
  actions: {
    setUserData({ commit }, payload) {
      commit('SET_USERUUID', payload.uuid)
      commit('SET_AVATARURL', payload.avatarURL)
      commit('SET_PHONE', payload.phone)
      commit('SET_ROLE', payload.role)
      commit('SET_USER_EXTEND', payload.userExtend)
      commit('SET_STATUS', payload.status)
    },
    clearUserData({ commit }) {
      commit('SET_AVATARURL', null)
      commit('SET_PHONE', null)
      commit('SET_ROLE', null)
      commit('SET_USER_EXTEND', null)
      commit('SET_STATUS', null)
    },
    logout({ commit }, callback) {
      commit('LOGOUT', callback)
    }
  }
};
