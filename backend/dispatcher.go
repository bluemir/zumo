package backend

import "github.com/sirupsen/logrus"

// this is main loop
func (b *backend) runDispatcher(events *SystemEvents) {
	for {
		select {
		case evt := <-events.Join:
			logrus.Debugf("[mainloop:Join] %+v", evt)
			b.userAgentManager.DeliverJoin(evt.UserName, evt.ChannelID)
		case evt := <-events.Leave:
			logrus.Debugf("[mainloop:Leave] %+v", evt)
			b.userAgentManager.DeliverLeave(evt.UserName, evt.ChannelID)
		case evt := <-events.CreateChannel:
			logrus.Debugf("[mainloop:CreateChannel] %+v", evt)
			// skip
		case evt := <-events.UpdateChannel:
			logrus.Debugf("[mainloop:UpdateChannel] %+v", evt)
			// skip
		case evt := <-events.DeleteChannel:
			logrus.Debugf("[mainloop:DeleteChannel] %+v", evt)
			// skip
		case err := <-events.Error:
			// just leave log...
			logrus.Errorf("[mainloop:Error] %s", err.Error())
		}
	}
}
