import akka.actor.{ ActorSystem, Props }
import akka.io.IO
import akka.util.Timeout
import akka.pattern.ask

import spray.can.Http

import scala.concurrent.duration._

object Main extends App {
  implicit val system = ActorSystem("My-System")
  val apiActor = system.actorOf(Props[ApiActor])

  implicit val timeout = Timeout(3 seconds)

  IO(Http) ? Http.Bind(apiActor, interface = "localhost", port = 8080)
}
