export default {
  namespaced: true,
  state: {
    status: null,
    uuid: null,
    avatarURL: null,
    phone: null,
    role: null,
    profile: null,
  },
  mutations: {
    SET_USERUUID: (state, payload) => state.uuid = payload,
    SET_AVATARURL: (state, payload) => state.avatarURL = payload,
    SET_PHONE: (state, payload) => state.phone = payload,
    SET_ROLE: (state, payload) => state.role = payload,
    SET_PROFILE: (state, payload) => state.profile = payload,
    SET_STATUS: (state, payload) => state.status = payload,
  },
  actions: {
    setUserData({ commit }, payload) {
      commit('SET_USERUUID', payload.uuid)
      commit('SET_AVATARURL', payload.avatarURL)
      commit('SET_PHONE', payload.phone)
      commit('SET_ROLE', payload.role)
      commit('SET_PROFILE', payload.userExtend)
      commit('SET_STATUS', payload.status)
    }
  }
};
