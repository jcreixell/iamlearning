import spray.routing._

trait ApiRoute extends HttpService {
  val myApiRoute =
    get {
      complete("Hello World!")
    }
}
