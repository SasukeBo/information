import Vue from 'vue';

import {
  Form,
  FormItem,
  Button,
  Input,
  Checkbox,
  Message,
  MessageBox,
  Row,
  Col,
  Tag,
  Drawer,
  CheckboxGroup,
  CheckboxButton,
  Select,
  Option,
  Loading,
  Pagination,
  Autocomplete,
  Transfer,
  Tooltip,
  DatePicker,
  Carousel,
  CarouselItem,
  InputNumber, Table, TableColumn
} from 'element-ui'

Vue.use(Table);
Vue.use(TableColumn);
Vue.use(InputNumber);
Vue.use(CarouselItem);
Vue.use(Carousel);
Vue.use(DatePicker);
Vue.use(Tooltip);
Vue.use(Autocomplete);
Vue.use(Transfer);
Vue.use(Pagination);
Vue.use(Form);
Vue.use(FormItem);
Vue.use(Button);
Vue.use(Input);
Vue.use(Checkbox);
Vue.use(Row);
Vue.use(Col);
Vue.use(Tag);
Vue.use(Drawer);
Vue.use(CheckboxGroup);
Vue.use(CheckboxButton);
Vue.use(Select);
Vue.use(Option);
Vue.use(Loading.directive);
Vue.prototype.$message = Message;
Vue.prototype.$alert = MessageBox.alert;
Vue.prototype.$confirm = MessageBox.confirm;
