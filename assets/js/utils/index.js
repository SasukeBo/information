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

function timeFormatter(timeStr, format = '%y年%m月%d日 %timestring') {
  if (!timeStr) {
    return '-';
  }

  var time = new Date(timeStr);
  var y = time.getFullYear();
  var m = time.getMonth() + 1;
  var d = time.getDate();
  var timeString = time.toTimeString().slice(0, 8);
  format = format.replace('%y', y);
  format = format.replace('%m', m);
  format = format.replace('%d', d);
  format = format.replace('%timestring', timeString);
  return format
}

export {
  timeFormatter,
  parseGQLError,
  getWeekDay,
  parseUserAgent
}
