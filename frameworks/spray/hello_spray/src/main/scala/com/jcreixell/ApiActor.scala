import akka.actor.Actor
import spray.routing._

class ApiActor extends Actor with ApiRoute {
  def actorRefFactory = context

  def receive = runRoute(myApiRoute)
}
