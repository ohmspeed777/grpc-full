import { FoodsServiceClient } from "../protos/FoodsServiceClientPb"
import configs from "../configs/configs"

const instants = {
  foodGRPC: new FoodsServiceClient(configs.ENVOY)
}

export default instants