function parseGQLError(e) {
  return JSON.parse(e.graphQLErrors[0].message)
}

function getWeekDay(date) {
  var weeks = ['星期日', '星期一', '星期二', '星期三', '星期四', '星期五', '星期六']
  return weeks[date.getDay()]
}

function parseUserAgent(ua) {
  var reg = new RegExp()
  reg.compile('[\\w/\\d\\.\\s]*\\(([^\\(\\)]+)\\) .*')
  var matches = reg.exec(ua)
  if (matches && matches.length > 1) return matches[1]

  return '未知设备'
}

function timeFormatter(timeStr) {
  var time = new Date(timeStr);
  var y = time.getFullYear();
  var month = time.getMonth() + 1;
  var day = time.getDate();
  var timeString = time.toTimeString().slice(0, 8);
  return `${y}年${month}月${day}日 ${timeString}`;
}

export {
  timeFormatter,
  parseGQLError,
  getWeekDay,
  parseUserAgent
}
