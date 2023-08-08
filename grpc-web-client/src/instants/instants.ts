import { FoodsServiceClient } from "../protos/FoodsServiceClientPb";
import { UserServiceClient } from "../protos/UsersServiceClientPb";
import configs from "../configs/configs";

const instants = {
  foodGRPC: new FoodsServiceClient(`${configs.ENVOY}`),
  userGRPC: new UserServiceClient(configs.ENVOY)
};

export default instants;
